//
//  ContentView.swift
//  Homie
//
//  Created by Mixko on 10/5/22.
//

import SwiftUI

struct ContentView: View {
    var body: some View {
        TabView {
            States()
                .tabItem {
                    Image(systemName: "list.dash")
                    Text("States")
                }
         
            Groups()
                .tabItem {
                    Image(systemName: "person.3.fill")
                    Text("Groups")
                }
         
            Accessories()
                .font(.system(size: 30, weight: .bold, design: .rounded))
                .tabItem {
                    Image(systemName: "server.rack")
                    Text("Accessories")
                }
         
            Settings()
                .font(.system(size: 30, weight: .bold, design: .rounded))
                .tabItem {
                    Image(systemName: "person.crop.circle")
                    Text("Profile")
                }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
