//
//  Groups.swift
//  Homie
//
//  Created by Mixko on 10/5/22.
//

import SwiftUI

struct Groups: View {
    var body: some View {
        NavigationView {
            List {
                ForEach (0..<20) {_ in
                    NavigationLink(destination: GroupNav()) {
                        GroupList().padding(.vertical,7)
                    }
                }
            }.navigationTitle("Groups")
        }
    }
}

struct Groups_Previews: PreviewProvider {
    static var previews: some View {
        Groups()
    }
}
