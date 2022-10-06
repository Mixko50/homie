//
//  GroupContainer.swift
//  Homie
//
//  Created by Mixko on 10/6/22.
//

import SwiftUI

struct GroupContainer: View {
    var body: some View {
        VStack {
            HStack {
                Image(systemName: "person.3.fill")
                Text("Latest created group").font(.system(size: 18, weight: .medium, design: .rounded))
            }.foregroundColor(Color.blue)
        }
    }
}

struct GroupContainer_Previews: PreviewProvider {
    static var previews: some View {
        GroupContainer()
    }
}
